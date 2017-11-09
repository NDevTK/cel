// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"fmt"
	"strings"
)

// IsRFC1035Label returns nil if |s| is a <label> as defined in RFC1035. If
// not, returns an error that's probably helpful.
//
// The RFC 1035 label production is:
//
//     <label>       ::= <letter> [ [ <ldh-str> ] <let-dig> ]
//
//     <ldh-str>     ::= <let-dig-hyp> | <let-dig-hyp> <ldh-str>
//
//     <let-dig-hyp> ::= <let-dig> | "-"
//
//     <let-dig>     ::= <letter> | <digit>
//
//     <letter>      ::= any one of the 52 alphabetic characters A through Z in
//                       upper case and a through z in lower case
//
//     <digit>       ::= any one of the ten digits 0 through 9
//
func IsRFC1035Label(s string) error {
	return isLabelWithOffset(s, s, 0)
}

func isLabelWithOffset(s, full_s string, offset int) error {
	var c byte
	if s == "" {
		if full_s != "" {
			return fmt.Errorf("%s is not a valid domain. Valid domains can't have empty subdomains", full_s)
		} else {
			return fmt.Errorf("empty string is not a valid RFC 1035 label. Valid labels can't be empty.")
		}
	}

	if len(s) > 63 {
		return fmt.Errorf("Valid RFC 1035 labels can't be longer than 63 characters.")
	}

	for i := 0; i < len(s); i++ {
		c = s[i]
		switch {
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
			// TODO(asanka): Reconsider whether we allow uppercase.
		case ('0' <= c && c <= '9') && i != 0:
		case c == '-' && i != 0:
		default:
			// All of the above are empty cases which match valid characters. If
			// we get here, then something is wrong.
			return fmt.Errorf("\"%s\" is not a valid RFC 1035 label. "+
				"The character '%c' at position %d is not valid.",
				full_s, c, (i + 1 + offset))
		}
	}

	if c == '-' {
		if full_s != s {
			return fmt.Errorf("\"%s\" is not a valid RFC 1035 domain. Subdomains can't end with a '-'.", full_s)
		} else {
			return fmt.Errorf("\"%s\" is not a valid RFC 1035 label. Cannot end with a '-'.", s)
		}
	}

	return nil
}

// IsRFC1035Domain return nil if |s| is a valid <subdomain> as defined in
// RFC1035. If not returns an error that's probably helpful.
//
// The RFC 1035 <subdomain> production is:
//
//     <subdomain> ::= <label> | <subdomain> "." <label>
//
//     <label>       ::= <letter> [ [ <ldh-str> ] <let-dig> ]
//
//     <ldh-str>     ::= <let-dig-hyp> | <let-dig-hyp> <ldh-str>
//
//     <let-dig-hyp> ::= <let-dig> | "-"
//
//     <let-dig>     ::= <letter> | <digit>
//
//     <letter>      ::= any one of the 52 alphabetic characters A through Z in
//                       upper case and a through z in lower case
//
//     <digit>       ::= any one of the ten digits 0 through 9
//
func IsRFC1035Domain(s string) error {
	if len(s) > 255 {
		return fmt.Errorf("domain names cannot be longer than 255 characters")
	}
	subdomains := strings.Split(s, ".")
	offset := 0
	for _, sub := range subdomains {
		if err := isLabelWithOffset(sub, s, offset); err != nil {
			return err
		}
		offset += len(sub) + 1
	}
	return nil
}

// IsRFC1035DomainLabel returns nil if |s| is matches the production "[
// <subdomain> ':' ] <label>".
//
// See IsRFC1035Label() and IsRFC1035Domain() for details on the <subdomain>
// and <label> productions.
func IsRFC1035DomainLabel(s string) error {
	if strings.ContainsRune(s, ':') {
		sp := strings.SplitN(s, ":", 2)
		err := IsRFC1035Domain(sp[0])
		if err != nil {
			return err
		}
		return IsRFC1035Label(sp[1])
	}
	return IsRFC1035Label(s)
}

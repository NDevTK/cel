#!/usr/bin/env python

# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from __future__ import print_function

import difflib
import logging
import os
import re
import textwrap


def ProcessIncludesInContent(lines, fname):
  '''ProcessIncludesInContent expands any INCLUDE directives found in a markdown file.

  Parameters:
    lines : a list of strings with line endings that represent the contents
            of the Markdown file.
    fname : the name of the file including the path if applicable. Includes
            are resolved relative to this path.

  Include directives take the form:

      <!-- INCLUDE <relative_path> [(<line_count> lines)] [fenced as <fence_type>] -->

  The parts of the directive in brackets are optional.

  The <relative_path> is an unquoted string with no spaces in it that
  specifies the relative path to the markdown file to be included. The lack
  of spaces in the paths can be easily remedied, but is not expected to be an
  issue for our use case in CEL.

  <fence_type> can be used to indicate that the content of the file should be
  included as a fenced block of the specified type. Note that the type is
  mandatory for the fence specification. If specified, the contents of the
  file will be included in the file as follows:

      ``` <fence_type>
      <contents of the file>
      ```

  Once expanded, the "(<line_count> lines)" portion will be inserted to
  indicate how many lines were added during the expansion. Do not change the
  line count or the included contents.
  '''

  class Include(object):

    def __init__(self, fn, lc, at, ft):
      self.filename = fn
      self.line_count = int(lc) if lc else 0
      self.at = at
      self.fence_type = ft

  dirname = os.path.dirname(fname)
  replacements = []

  for i, l in enumerate(lines):
    if l.startswith('<!-- INCLUDE '):
      m = re.match(
          r'<!-- INCLUDE +(?P<fn>[^ ]*) +(?:\((?P<lc>\d+) lines\) +|)(?:fenced as (?P<ft>\w+) +|)-->',
          l)
      if m is None:
        continue
      replacements.append(
          Include(m.group('fn'), m.group('lc'), i, m.group('ft')))

  for r in reversed(replacements):
    newlines = []
    with open(os.path.join(dirname, r.filename), 'r') as f:
      newlines = f.readlines()

    for l in newlines:
      if l.startswith('<!-- INCLUDE '):
        raise Exception('''Recursive includes are not supported.''')

    if r.fence_type is not None:
      newlines[0:0] = ['``` {}\n'.format(r.fence_type)]
      newlines.append('```\n')

    lines[r.at] = '<!-- INCLUDE {} ({} lines){} -->\n'.format(
        r.filename, len(newlines), ' fenced as {}'.format(r.fence_type)
        if r.fence_type else '')
    lines[r.at + 1:r.at + 1 + r.line_count] = newlines


def FixOldStyleLinks(lines, fname):
  '''FixOldStyleLinks replaces links of the form [foo] with [foo][]
  
  The former style is accepted by Gitliles, but is not valid CommonMark. Hence
  this function replaces it with the equivalent latter form. This replacement
  ensures that the links are correctly handled by editors and viewers other
  than Gitiles.
  '''

  bad_link_re = re.compile(r'(^|[^]])\[(?P<ref>[^]]+)\](?=$|[^[:(])')
  in_pre = False
  for i, l in enumerate(lines):
    if '```' in l:
      in_pre = not in_pre

    if in_pre:
      continue

    if "[TOC]" in l:
      continue

    lines[i] = re.sub(bad_link_re, r'\1[\g<ref>][]', l)


def FixTrailingWhitespace(lines, fname):
  '''FixTrailingWhitespace does what it says and removes trailing whitespace.
  '''
  trailing_ws_re = re.compile(r'\s+(?=\n)$')
  for i, l in enumerate(lines):
    lines[i] = re.sub(trailing_ws_re, '', l)


def CheckLinksInContent(lines, fname):
  '''CheckLinksInContent verifies that reference style links are defined in the
  same document.

  Referene style links are links of the form [foo][], or [Foo][foo] where [foo]
  needs to be defiend somewhere else in the document as:

      [foo]: https://example.com/foo

  This function raises an exception with a suitable description if a reference
  style link is not defined.
  '''

  links = set()
  link_re = re.compile(r"\[(?P<ref>[^]]+)\]: ")
  for l in lines:
    m = link_re.match(l)
    if m is None:
      continue
    links.add(m.group('ref'))

  whole_thing = re.sub(r'\s+', ' ', ''.join(lines), count=0)
  not_found = set()
  for m in re.finditer(r'\[(?P<a>[^]]*)\]\[(?P<ref>[^]]*)\]', whole_thing):
    ref = m.group('ref') if m.group('ref') != '' else m.group('a')
    if ref not in links:
      not_found.add(ref)

  if len(not_found) != 0:
    raise Exception(
        textwrap.dedent('''\
            The following list of links were unresolved in {}:
                {}
            '''.format(fname, ','.join(list(not_found)))))


def FormatMarkdown(fname, dry_run=False):
  '''FormatMarkdown resolves any includes, fixes links, corrects trailing
  whitespace, and verifies that reference style links are defined in Markdown
  document specified by the filename in |fname|.

  if |dry_run| is True, then checks whether the contents in the file at |fname|
  would be modified by the function.

  Returns True if the file at |fname| was modified (or in the case of
  |dry_run==True|, would be modified). False otherwise. A return value of False
  can be safely assumed to mean that the file contents were not modified.

  For additional information about the changes that are applied see the
  documentation in:

      ProcessIncludesInContent()
      FixOldStyleLinks()
      FixTrailingWhitespace()
      CheckLinksInContent()
  '''

  lines = []
  with open(fname, 'r') as f:
    lines = f.readlines()
  unmodified = lines[:]

  ProcessIncludesInContent(lines, fname)
  FixOldStyleLinks(lines, fname)
  FixTrailingWhitespace(lines, fname)
  CheckLinksInContent(lines, fname)

  if lines == unmodified:
    logging.info("%s is already correctly formatted", fname)
    return False

  if dry_run:
    print("Would write %s with the following changes:", fname)
    for l in difflib.unified_diff(unmodified, lines, fname + "  (original)",
                                  fname + "  (modified)"):
      print(l, end='')
    print('\n')
  else:
    logging.info("Writing %s", fname)
    with open(fname, 'w') as f:
      f.writelines(lines)

  return True

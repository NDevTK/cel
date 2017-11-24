#!/usr/bin/env python

# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import unittest
import os
import sys
import textwrap

BUILD_PATH = os.path.dirname(os.path.realpath(__file__))
TESTDATA_PATH = os.path.join(BUILD_PATH, 'testdata')

sys.path.append(BUILD_PATH)

from markdown_utils import FormatMarkdown, ProcessIncludesInContent, FixOldStyleLinks, FixTrailingWhitespace, CheckLinksInContent


class TestMarkdownUtils(unittest.TestCase):

  def test_includes_single(self):
    lines = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE included.md -->
                Postlude
                ''').splitlines(True)
    fname = os.path.join(TESTDATA_PATH, 'foo.md')

    ProcessIncludesInContent(lines, fname)

    expected = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE included.md (2 lines) -->
                Included stuff.

                Postlude
                ''').splitlines(True)

    self.assertListEqual(lines, expected)

    # ProcessIncludesInContent should be idempotent.
    ProcessIncludesInContent(lines, fname)
    self.assertListEqual(lines, expected)

  def test_includes_single_relative(self):
    lines = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE testdata/included.md -->
                Postlude
                ''').splitlines(True)
    fname = os.path.join(BUILD_PATH, 'foo.md')

    ProcessIncludesInContent(lines, fname)

    expected = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE testdata/included.md (2 lines) -->
                Included stuff.

                Postlude
                ''').splitlines(True)

    self.assertListEqual(lines, expected)

  def test_includes_whitespace(self):
    lines = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE      testdata/included.md   (0 lines)   fenced as foo     -->
                Postlude
                ''').splitlines(True)
    fname = os.path.join(BUILD_PATH, 'foo.md')

    ProcessIncludesInContent(lines, fname)

    expected = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE testdata/included.md (4 lines) fenced as foo -->
                ``` foo
                Included stuff.

                ```
                Postlude
                ''').splitlines(True)

    self.assertListEqual(lines, expected)

  def test_includes_fenced(self):
    lines = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE included.md fenced as text -->
                Postlude
                ''').splitlines(True)
    fname = os.path.join(TESTDATA_PATH, 'foo.md')

    ProcessIncludesInContent(lines, fname)

    expected = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE included.md (4 lines) fenced as text -->
                ``` text
                Included stuff.

                ```
                Postlude
                ''').splitlines(True)

    self.assertListEqual(lines, expected)

    # ProcessIncludesInContent should be idempotent.
    ProcessIncludesInContent(lines, fname)
    self.assertListEqual(lines, expected)

  def test_includes_multiple(self):
    lines = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE included.md -->
                <!-- INCLUDE included.md -->
                <!-- INCLUDE included.md -->
                Postlude
                ''').splitlines(True)
    fname = os.path.join(TESTDATA_PATH, 'foo.md')

    ProcessIncludesInContent(lines, fname)

    expected = textwrap.dedent('''\
                Prelude
                <!-- INCLUDE included.md (2 lines) -->
                Included stuff.

                <!-- INCLUDE included.md (2 lines) -->
                Included stuff.

                <!-- INCLUDE included.md (2 lines) -->
                Included stuff.

                Postlude
                ''').splitlines(True)

    self.assertListEqual(lines, expected)

    # ProcessIncludesInContent should be idempotent.
    ProcessIncludesInContent(lines, fname)
    self.assertListEqual(lines, expected)

  def test_old_style_links(self):
    lines = textwrap.dedent('''\
                Some text
                foo is a [foo] and must be bar.
                foo is a [foo] and must be [bar].

                [TOC]

                why is [toc] special?
                ''').splitlines(True)
    fname = os.path.join(BUILD_PATH, 'foo.md')

    FixOldStyleLinks(lines, fname)

    expected = textwrap.dedent('''\
                Some text
                foo is a [foo][] and must be bar.
                foo is a [foo][] and must be [bar][].

                [TOC]

                why is [toc][] special?
                ''').splitlines(True)

    self.assertListEqual(lines, expected)

  def test_trailing_whitespace(self):
    lines = textwrap.dedent('''\
                Foo
                Bar
                ''').splitlines(True)
    fname = os.path.join(BUILD_PATH, 'foo.md')

    # So that presubmits don't complain.
    lines[1] = 'Bar   \n'

    FixTrailingWhitespace(lines, fname)

    expected = textwrap.dedent('''\
                Foo
                Bar
                ''').splitlines(True)
    self.assertListEqual(lines, expected)

  def test_links(self):
    fname = os.path.join(BUILD_PATH, 'foo.md')
    lines = textwrap.dedent('''\
                This is a [link][]

                [link]: https://example.com/link
                ''').splitlines(True)

    # This shouldn't raise any exceptions.
    CheckLinksInContent(lines, fname)

    lines = textwrap.dedent('''\
                This is a [link][] that is not defined anywhere
                so CheckLinksInContent will raise an exception.
                ''').splitlines(True)
    with self.assertRaisesRegexp(Exception, 'unresolved'):
      CheckLinksInContent(lines, fname)

    lines = textwrap.dedent('''\
                This is a [link]
                but is not identified as a links since it's not
                a valid CommonMark link.
                ''').splitlines(True)
    CheckLinksInContent(lines, fname)

    lines = textwrap.dedent('''\
                This is [link1][], and this [is][link2]
                which are two lines on the same line. The second link should be
                resolved by defining [link2].
                ''').splitlines(True)
    with self.assertRaisesRegexp(Exception, 'link1'):
      CheckLinksInContent(lines, fname)

    lines.append('[link1]: foo\n')
    # still raises since the second link is not defined.
    with self.assertRaisesRegexp(Exception, 'link2'):
      CheckLinksInContent(lines, fname)

    lines.append('[link2]: foo\n')
    CheckLinksInContent(lines, fname)


if __name__ == '__main__':
  unittest.main()

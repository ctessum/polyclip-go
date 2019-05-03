// Copyright (c) 2011 Mateusz Czapliński
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package polyclip

import (
	. "testing"
)

func TestChainLinkChain(t *T) {
	a := chain{points: []Point{{0, 1}, {0, 2}, {0, 3}, {1, 1}}}
	b := chain{points: []Point{{1, 1}, {1, 2}}}
	verify(t, a.linkChain(&b), "Expected being able to link chains")
	verify(t, len(a.points) == 5, "Expected len==5, got %d", len(a.points))
}

func TestChainLinkChainWithTolerance(t *T) {
	a := chain{points: []Point{{0, 1}, {0, 2}, {0, 3}, {1.000000000000001, 1.000000000000001}}}
	b := chain{points: []Point{{1, 1}, {1, 2}}}
	verify(t, a.linkChainWithTolerance(&b, 3e-14), "Expected being able to link chains")
	verify(t, len(a.points) == 5, "Expected len==5, got %d", len(a.points))
}

func TestChainCloseWithTolerance(t *T) {
	a := chain{points: []Point{{0, 1}, {0, 2}, {0, 3}, {0.000000000000001, 1.000000000000001}}}
	verify(t, a.closeWithTolerance(3e-14), "Expected being able to close the chains")
	verify(t, a.closed, "Expected chain to be closed.")
	verify(t, len(a.points) == 4, "Expected len==4, got %d", len(a.points)) // Note: points are connected, not merged.
}

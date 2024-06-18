/* This program is licensed under the MIT license:
 *
 * Copyright 2024 Simon de Vlieger
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */

package anteater

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/mmcdole/gofeed"
)

var feeds = []*Feed{
	&Feed{Nickname: "supakeen", URL: "https://supakeen.com/weblog/feeds/all.atom.xml"},
	&Feed{Nickname: "hansg", URL: "https://hansdegoede.dreamwidth.org/data/atom"},
	&Feed{Nickname: "eloy", URL: "https://eloydegen.com/blog/index.xml"},
	&Feed{Nickname: "shiz", URL: "https://shizmob.tumblr.com/rss"},
}

func Poll(state *State, c chan<- Item) error {
	feedParser := gofeed.NewParser()

	for _, feed := range feeds {
		feedReader, err := feedParser.ParseURL(feed.URL)

		if err != nil {
			slog.Info(fmt.Sprintf("poll: feed '%v' failed to parse due to: %v\n", feed.Nickname, err))
			continue
		}

		// XXX Do we want to use fd.PublishedParsed too?
		for _, item := range feedReader.Items {
			// We skip items if we can't parse their published date.
			if item.PublishedParsed == nil {
				slog.Info(fmt.Sprintf("poll: couldn't parse the `published` field for feed '%v'\n", feed.Nickname))
				continue
			}

			// Otherwise we'll need a title and a link
			if len(item.Title) == 0 {
				slog.Info(fmt.Sprintf("poll: couldn't parse the `time` field for feed '%v'\n", feed.Nickname))
				continue
			}

			if len(item.Link) == 0 {
				slog.Info(fmt.Sprintf("poll: couldn't parse the `link` field for feed '%v'\n", feed.Nickname))
				continue
			}

			// See if the item is newer than the last time we polled, if the program starts up the LastPolledAt time is
			// set to the program startup time.
			if state.LastPolledAt.After(*item.PublishedParsed) {
				slog.Info(fmt.Sprintf("poll: ignoring stale item '%v' for feed '%v'\n", item.Title, feed.Nickname))
				continue
			}

			slog.Info(fmt.Sprintf("poll: submitting item '%v' for feed '%v'\n", item.Title, feed.Nickname))

			c <- NewItem(feed, item.Link)
		}
	}

	state.LastPolledAt = time.Now()

	return nil
}

func Loop(every time.Duration, state *State, c chan<- Item) error {
	for {
		Poll(state, c)
		time.Sleep(every)
	}
}

// SPDX-License-Identifier: MIT
// vim: ts=4 sw=4 noet

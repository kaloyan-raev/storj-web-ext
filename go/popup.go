// Copyright (C) 2018 Kaloyan Raev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"fmt"

	"honnef.co/go/js/dom"
	"storj.io/storj"
)

var (
	doc = dom.GetWindow().Document()
)

func renderStatus(statusText string) {
	doc.GetElementByID("status").SetInnerHTML(statusText)
}

func main() {
	env := storj.Env{
		URL: storj.DefaultURL,
	}

	// TODO make this call async
	info, err := storj.GetInfo(env)
	if err != nil {
		renderStatus(fmt.Sprintf("%v", err))
	} else {
		renderStatus(fmt.Sprintf("Title: %s<br>Description: %s<br>Version: %s<br>Host: %s\n",
			info.Title, info.Description, info.Version, info.Host))
	}
}

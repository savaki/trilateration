//	Copyright 2015 Matt Ho
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.

package trilateration

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCenter(t *testing.T) {
	Convey("Given a set of points", t, func() {
		r := 25.0

		p1 := Point{
			R: r,
		}
		p2 := Point{
			X: r * .8,
			R: r,
		}
		p3 := Point{
			Y: r * .8,
			R: r,
		}

		solution, err := Solve(p1, p2, p3)
		So(err, ShouldBeNil)
		So(len(solution), ShouldEqual, 3)

		So(solution.First(), ShouldResemble, solution[0])
		So(solution[0].X, ShouldEqual, 10)
		So(solution[0].Y, ShouldEqual, 10)
		So(solution[1].X, ShouldEqual, 0)
		So(solution[1].Y, ShouldEqual, 20)
		So(solution[2].X, ShouldEqual, 0)
		So(solution[2].Y, ShouldEqual, 20)
	})
}

package volumes

import (
  "fmt"
  "testing"
)

func TestMortonFunctions(t *testing.T) {
  xyz := make([]int64, 3)
  var morton int64

  // test xyz -> morton index
  // stage 1
  xyz[0] = 0
  xyz[1] = 0
  xyz[2] = 0
  morton = XYZMorton(xyz)
  if morton != int64(0) {
    t.Error(fmt.Sprintf("Error: Morton test failed at stage 1. Could not convert (%d,%d,%d) to 0 (got morton index %d instead)", xyz[0], xyz[1], xyz[2], morton))
  }

  // stage 2
  xyz[0] = 1
  xyz[1] = 1
  xyz[2] = 1
  morton = XYZMorton(xyz)
  if morton != int64(7) {
    t.Error(fmt.Sprintf("Error: Morton test failed at stage 2. Could not convert (%d,%d,%d) to 7 (got morton index %d instead)", xyz[0], xyz[1], xyz[2], morton))
  }

  // stage 3
  xyz[0] = 1
  xyz[1] = 2
  xyz[2] = 3
  morton = XYZMorton(xyz)
  if morton != int64(53) {
    t.Error(fmt.Sprintf("Error: Morton test failed at stage 1. Could not convert (%d,%d,%d) to 53 (got morton index %d instead)", xyz[0], xyz[1], xyz[2], morton))
  }

  // stage 4
  xyz[0] = 500
  xyz[1] = 800
  xyz[2] = 200
  morton = XYZMorton(xyz)
  if morton != int64(330668096) {
    t.Error(fmt.Sprintf("Error: Morton test failed at stage 4. Could not convert (%d,%d,%d) to 330668096 (got morton index %d instead)", xyz[0], xyz[1], xyz[2], morton))
  }

  // test morton index to xyz
  // stage 5
  morton = int64(330668096)
  xyz = MortonXYZ(morton)
  if xyz[0] != 500 && xyz[1] != 800 && xyz[2] != 200 {
      t.Error(fmt.Sprintf("Error: Morton test failed at stage 5. Could not convert %d to (500,800,200) (got %d,%d,%d instead)", morton, xyz[0], xyz[1], xyz[2]))
  }

  // stage 6
  morton = int64(0)
  xyz = MortonXYZ(morton)
  if xyz[0] != 0 && xyz[1] != 0 && xyz[2] != 0 {
      t.Error(fmt.Sprintf("Error: Morton test failed at stage 5. Could not convert %d to (0,0,0) (got %d,%d,%d instead)", morton, xyz[0], xyz[1], xyz[2]))
  }

  // stage 7
  morton = int64(50)
  xyz = MortonXYZ(morton)
  if xyz[0] != 0 && xyz[1] != 3 && xyz[2] != 2 {
      t.Error(fmt.Sprintf("Error: Morton test failed at stage 5. Could not convert %d to (0,3,2) (got %d,%d,%d instead)", morton, xyz[0], xyz[1], xyz[2]))
  }

}

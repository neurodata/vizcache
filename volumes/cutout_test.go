package volumes

import (
  "testing"
)

func TestGet (t* testing.T) {
  vc := NewVolumeCache ( [3]uint{512,512,128}, "test" )
  vc.Get ( 0,1000,0,2000,0,3000 )
}

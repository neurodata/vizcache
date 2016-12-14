package volumes

/* XYZMorton generates a morton order index from integer XYZ coordinates */
func XYZMorton(xyz []int64) int64 {
  var mask, morton int64
  mask = 0x001

  x := xyz[0]
  y := xyz[1]
  z := xyz[2]

  for i := uint32(0); i < 21; i++ {
    morton += ( x & mask ) << (2*i)
    morton += ( y & mask ) << (2*i+1)
    morton += ( z & mask ) << (2*i+2)

    mask = mask << 1
  }

  return morton
}

/* MortonXYZ generates integer XYZ coordinates from a morton order index */
func MortonXYZ(morton int64) []int64 {
  var xmask, ymask, zmask int64
  xmask = 0x001
  ymask = 0x002
  zmask = 0x004

  output := make([]int64, 3)

  for i := uint32(0); i < 21; i++ {
    output[0] = output[0] + ( ( xmask & morton ) << i )
    output[1] = output[1] + ( ( (ymask & morton) << i ) >> 1 )
    output[2] = output[2] + ( ( (zmask & morton) << i ) >> 2 )
    morton = morton >> 3
  }

  return output
}

package volumes

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "github.com/kowalczykp/go-blosc"
)

type VolumeCache struct
{
  Cubedim [3]uint
  Source string
}

// factory function to initialize a volume cache
func NewVolumeCache ( cubedim [3]uint, source string ) *VolumeCache {
  vc := new(VolumeCache)
  vc.Cubedim = cubedim
  vc.Source = source
  return vc
}

// main entry point.  Get a volume.
//  This will construct a volume out of the cache and fill in missing parts
//  by requesting them from a data source
//  This will retrieve the volume in aligned cuboids from a source.
func (vc* VolumeCache) Get ( xmin uint, xmax uint, ymin uint, ymax uint, zmin uint, zmax uint ) []byte {

  fmt.Println(`-> fmt.Println("In Volume Cache")`)

  // RBTODO put in class?
  //enc := blosc.NewEncoder()
  dec := blosc.NewDecoder()

  // variables for dimensions
  xcubedim := vc.Cubedim[0]
  ycubedim := vc.Cubedim[1]
  zcubedim := vc.Cubedim[2]

  // Geometry of the cutout in cuboids.  Start cube.
  zstartcube := zmin/zcubedim
  ystartcube := ymin/ycubedim
  xstartcube := xmin/xcubedim

  // Number of cubes
  znumcubes := (zmax+zcubedim-1)/zcubedim - zstartcube
  ynumcubes := (ymax+ycubedim-1)/ycubedim - ystartcube
  xnumcubes := (xmax+xcubedim-1)/xcubedim - xstartcube

  numels := znumcubes*ynumcubes*xnumcubes

  var arrofidxs = make([]int64,numels)

  // build the list of morton indexes needed 
  xyz := make ([]int64,3)
  var x,y,z,i uint
  i=0
  for z = 0; z < znumcubes; z++ {
    for y = 0; y < ynumcubes; y++ {
      for x = 0; x < xnumcubes; x++ {
        xyz[0] = int64(x)
        xyz[1] = int64(y)
        xyz[2] = int64(z)
        arrofidxs[i] = XYZMorton ( xyz )
        i = i+1
      }
    }
  }

  // You have a list of indexes. Check to see if each is in cache?  
  // if not fetch it from the server
  for i=0; i<numels; i++ {

    // godebug: break
    _ = "breakpoint"

    xyz := MortonXYZ ( arrofidxs[i] )
    xstart := uint(xyz[0])*xcubedim
    ystart := uint(xyz[1])*ycubedim
    zstart := uint(xyz[1])*zcubedim

    //RBTODO get this all from the VC source information
    url := fmt.Sprintf("http://brainviz1.cs.jhu.edu/microns/nd/sd/kasthuri11/image/blosc/%d/%d,%d/%d,%d/%d,%d/",1,xstart,xstart+xcubedim,ystart,ystart+ycubedim,zstart,zstart+zcubedim)

    fmt.Printf(url+"\n")
    resp, err := http.Get(url)
    if err != nil {
      panic(err)
    }

    // Read the blosc cube from the response
    blosccube, err := ioutil.ReadAll ( resp.Body )
    if err != nil {
      panic(err)
    }
    resp.Body.Close()

    // Extract the blosc cube
    cube, err := dec.Decode(blosccube, nil)
    if err != nil {
      panic(err)
    }
    fmt.Printf("%s\n",cube)
  }


  return []byte{0x00,0x01}
} 

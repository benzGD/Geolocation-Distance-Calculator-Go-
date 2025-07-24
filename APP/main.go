package main

import(
    "fmt"
    "geodistance/distance"
    "geodistance/locationmap"



)

func main() {
   earth := distance.World{Radius: 6371.0} 

   name1 := "Las Vegas, USA"
   name2:= "Area 51, USA"

    m1,v := locationmap.PlacesCoordinates[name1]
    if !v {
        fmt.Println("invalid request")
        return
    }
    m2,v0 := locationmap.PlacesCoordinates[name2]
    if !v0 {
        fmt.Println("invalid request")
        return
    }
    p1 :=distance.CoordinateToLocation(m1.Glat, m1.Glong)
    p2 :=distance.CoordinateToLocation(m2.Glat, m2.Glong)


    fmt.Printf("the distance between '%v' & '%v' is : %.4fkms",name1, name2,earth.Distance(p1,p2))

    


    // fmt.Printf("%+v",locationmap.PlacesCoordinates["Mount Everest, Nepal"])







   
//    mount_everest := distance.GeoCoordinate{Glat: distance.Coordinate{D: 27, M: 59, S: 17, H: 'N'}, Glong: distance.Coordinate{D: 86, M: 55, S: 31, H: 'E'}} 
//    Area_51 := distance.GeoCoordinate{Glat: distance.Coordinate{D: 37, M: 14, S: 36, H: 'N'}, Glong: distance.Coordinate{D: 115, M: 48, S: 40, H: 'W'}} 
   
// g21

//    fmt.Println("distance :", earth.Distance(m1,m2) )





}
    

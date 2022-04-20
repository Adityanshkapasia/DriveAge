import { useEffect, useState } from "react"
import internal from "stream"
import { serialize } from "v8"

type Allcars= {
    
    id:           number,
	brand:        string,
	type:         string,
	price:        number,
	MPG:          number,
	name:         string,
	tankcapacity: number,    
	color:        string  
        
}

export default function Viewallcars() {
    const [cars, setcars] = useState<Allcars[]>([])
    const onLoad=()=>{
        fetch("http://localhost:8080/allcar", {
            credentials : 'include' ,
        }) .then (res=>{
            if (res.status=== 200){
                res.json().then(t=>{
                    setcars(t)
                })
            }
            else {
               
            }
        })
    }
    useEffect(() => {
      onLoad()
    
    }, [])
    
    return(
        <div>
       {
           cars.map(car=>
               (
             <div key={car.id}>
                 {car.name} {car.brand} {car.color} {car.price} {car.tankcapacity}  {car.type} 

             </div>
             
               )
           )
       }
        </div>
    )
}
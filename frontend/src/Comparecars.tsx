import { useEffect, useState } from "react"

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



export default function(){
    const [car1, setCar1] = useState<Allcars>()
    const [car2, setCar2] = useState<Allcars>()
    const [cars, setCars] = useState<Allcars[]>([])

    const onLoad=()=>{
        fetch("http://localhost:8080/allcar", {
            credentials : 'include' ,
        }) .then (res=>{
            if (res.status=== 200){
                res.json().then(t=>{
                    setCars(t)
                })
            }
            else {
               
            }
        })
    }
    useEffect(() => {
      onLoad()
    
    }, [])
    

    const loadCar=(id: number, setcar:React.Dispatch<React.SetStateAction<Allcars | undefined>>)=>{
        fetch("http://localhost:8080/car/"+id, {
            credentials : 'include' ,
        }) .then (res=>{
            if (res.status=== 200){
                res.json().then(t=>{
                    setcar(t)
                })
            }
            else {
               
            }
        })
    }
    return(
    <div>
        <table>
  <tr>
    <th>
        <select>
            {cars.map(car => (
                <option value={car.id}>{car.name}</option>
            ))}
        </select>
    </th>
    <th>
        <select>
            {cars.map(car => (
                <option value={car.id}>{car.name}</option>
            ))}
        </select>
        </th>
    
  </tr>
  <tr>
    <td>Alfreds Futterkiste</td>
    <td>Maria Anders</td>
   
  </tr>
  <tr>
    <td>Centro comercial Moctezuma</td>
    <td>Francisco Chang</td>
  </tr>
</table>
    </div>
    )
}
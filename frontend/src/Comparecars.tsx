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
    

    const loadCar=(id: string, setcar:React.Dispatch<React.SetStateAction<Allcars | undefined>>)=>{
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
        <select onChange={e => {loadCar(e.target.value, setCar1)}}>
            <option selected disabled> Select Car</option>
            {cars.map(car => (
                <option value={car.id}>{car.name}</option>
            ))}
        </select>
    </th>
    <th>
        <select onChange={e => {loadCar(e.target.value, setCar2)}}>
            <option selected disabled> Select Car</option>
            {cars.map(car => (
                <option value={car.id}>{car.name}</option>
            ))}
        </select>
        </th>

  </tr>

          <tr>
              <td>{car1 && car1.name}</td>
              <td>{car2 && car2.name}</td>
          </tr>
          <tr>
              <td>{car1 && car1.brand}</td>
              <td>{car2 && car2.brand}</td>
          </tr>
          <tr>
              <td>{car1 && car1.color}</td>
              <td>{car2 && car2.color}</td>
          </tr>
          <tr>
              <td>{car1 && car1.price}</td>
              <td>{car2 && car2.price}</td>
          </tr>
          <tr>
              <td>{car1 && car1.tankcapacity}</td>
              <td>{car2 && car2.tankcapacity}</td>
          </tr>
          <tr>
              <td>{car1 && car1.type}</td>
              <td>{car2 && car2.type}</td>
          </tr>


</table>
    </div>
    )
}
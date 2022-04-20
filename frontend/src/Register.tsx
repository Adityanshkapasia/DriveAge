import { useState } from "react"
import { useNavigate } from "react-router-dom"

export default function Register(){

const [email, setemail] = useState("")
const [password, setpassword] = useState("")
const [name, setname] = useState("")
const navigate= useNavigate()
const onSubmit=()=>{
    fetch("http://localhost:8080/register", {
        method: 'POST',
        credentials : 'include' ,
        body: JSON.stringify({    "email": email,
        "password": password,
        "name": name})
    }) .then (res=>{
        if (res.status=== 200){
            alert("Successfully Registered!!!!!!!")
            navigate("/Login")
         
        }
        else {
           res.text().then(t=>{
               alert(t)
           })
        }
    })
}
return (
    <div> 
        email
        <input type="email" value={email} onChange={e=>{setemail(e.target.value)}} />

        <br />
        password
        <input type="password" value={password} onChange={e=>{setpassword(e.target.value)}} />
        <br />
        name
        <input type="name" value={name} onChange={e=>{setname(e.target.value)}} />
        <br />

        <button onClick={e=> onSubmit()}>
            Register
           
        </button>
    </div>
)
}
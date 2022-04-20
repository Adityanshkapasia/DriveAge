import { useState } from "react"
import { useNavigate } from "react-router-dom"

export default function Createblogpost(){
   const [title, settitle] = useState("")
   const [body, setBody] = useState("")
   const navigate=useNavigate()
       const onSubmit=()=>{
        fetch("http://localhost:8080/newpost", {
            method: 'POST',
            credentials: 'include',
            body : JSON.stringify({"title":title,
            "desc":body})
        }).then(res=>{
            if (res.status===200){
                alert("Successfullly Created Blog Post!!!!!")
                navigate("/blog")
            }
            else{
               res.text().then(t=>{
                   alert(t)
               })
            }
        })
    }
    return(
        <div>
        Title
        <br></br>
     <input type={"text"} value={title} onChange={e=>{settitle(e.target.value)}}></input>
     <br></br>
        Body
        <br></br>
     <input type={"text"} value={body} onChange={e=>{setBody(e.target.value)}}></input>
     <br></br>
     <button onClick={e=>onSubmit()}>Create</button>
        </div>
    )
}
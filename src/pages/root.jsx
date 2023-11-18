
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
const Root = () =>{
    const navigate = useNavigate();
    return(
        <div>

            <button onClick={()=>{navigate("/home")}}>Go Home</button>
        </div>
    )
}

export default Root
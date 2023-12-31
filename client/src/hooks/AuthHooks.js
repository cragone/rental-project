import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

axios.defaults.withCredentials = true;

const useSession = () => {

    const navigate = useNavigate()
    const [urlRoot, setUrlRoot] = useState("http://localhost")
    const [user, setUser] = useState("")

    useEffect(()=>{
        axios.get(urlRoot+"/auth/session").then((response)=>{
            setUser(response.data.response)
        }).catch((error)=>{
            navigate("/")
        })
    },[])

    return user
}
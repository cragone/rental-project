import React, { useState, useEffect } from 'react'
import axios from 'axios'

axios.defaults.withCredentials = true;

const usePropertyList = () => {

    const [urlRoot, setUrlRoot] = useState("http://localhost")
    const [propertyInfo, setPropertyInfo] = useState([])
    useEffect(()=>{
        axios.get(urlRoot+"/property/list").then((response)=>{
            setPropertyInfo(response.data.propertyList)
        }).catch((error)=>{
        })
    },[])

    return propertyInfo
}

const usePropertyTennantList = (props) =>{
    const [urlRoot, setUrlRoot] = useState("http://localhost")
    const [tennantList, setTennantList] = useState([])

    useEffect(()=>{
        const payload = {
            propertyID: props
        }
        axios.post(urlRoot+"/tennant/property_tennants", payload).then((response)=>{
            setTennantList(response.data.response)
        }).catch((error)=>{
        })
    },[])

    return tennantList
}

const tennantInfo = (props) => {
    const urlRoot = "http://localhost"


    const payload = {
        tennantID: props
    }
    const value = axios.post(urlRoot+"/tennant/get", payload).then((response)=>{
        return (response.data.tennant)
    }).catch((error)=>{
        return null
    })

    return value
}

export {usePropertyList, usePropertyTennantList, tennantInfo}
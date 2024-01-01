import React, { useState, useEffect } from 'react'
import axios from 'axios'

axios.defaults.withCredentials = true;

const usePropertyList = () => {

    const [urlRoot, setUrlRoot] = useState("http://localhost")
    const [propertyInfo, setPropertyInfo] = useState([])

    useEffect(()=>{
        axios.get(urlRoot+"/property/list").then((response)=>{
            console.log(response.data.propertyList)
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
            console.log(response.data.response)
            setTennantList(response.data.response)
        }).catch((error)=>{
        })
    },[])

    return tennantList
}

export {usePropertyList, usePropertyTennantList}
import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

const PaymentTest = () => {

    const [api, setApi] = useState(import.meta.env.VITE_API_PATH)

    const [order, setOrder] = useState("")

    // WORKFLOW:
    // Check order stats and if not filled get the id:

    const newOrder = () => {
        const payload = {
            invoiceID: "b8a2572c-f026-474c-bf3b-9cbe71ad5880"
        }
        axios.post(`http://localhost/invoice/create_order`, payload).then((response) => {
            console.log(response.data.response)
            setOrder(response.data.response)
        }).catch((error) => {
            console.log("error")
            console.log(error.response.data.error)
        })
    }

    const openPayment = () => {
        window.open(`https://www.sandbox.paypal.com/checkoutnow?token=${order}`, '_blank');
    }

    // &fundingSource=venmo

    const openVenmoPayment = () => {
        window.open(`https://www.sandbox.paypal.com/checkoutnow?token=${order}&fundingSource=venmo`, '_blank');
    }

    return (

        <>
            <button onClick={newOrder}>Create Order</button>

            <button onClick={openPayment}>Open Paypal Payment</button>
            <button onClick={openVenmoPayment}>Open Venmo Payment</button>


        </>
    )
}

export default PaymentTest
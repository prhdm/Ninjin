import React, { useState } from "react"
import NavbarComp from "../../components/Navbar/NavbarComp";
import { InputGroup, FormControl, Button} from 'react-bootstrap';
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from "axios";
import useToken from "../../components/App/useToken";

const Water = () => {
    const [deviceSerial, setSerial] = useState('')
    const [waterAmount, setAmount] = useState('')

    const url_addWateringLog = "http://usagi.carriot.ir:8000/watering-log/create";
    const token = useToken().token;

    const addNewWateringLog = (device_serial, waterAmount) =>{
        const wateringLogData = JSON.stringify({
            "time": "2008-10-19T10:23:54+02:00",
            "device_serial" : device_serial,
            "water_amount" : waterAmount,
        });

        axios.post(url_addWateringLog, wateringLogData , {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
            .then((response) => {
                console.log(response)
                window.location.reload(true)
            });


    }

    return (
        <>
            <NavbarComp />
            <span className="gap"></span>

            <div className="page-title">
                تنظیم مقدار آبدهی
            </div>

            <InputGroup dir="rtl" className="mb-3 humidity-input" size="lg">
                <FormControl aria-label="deviceSerial" placeholder="سریال دستگاه"
                             onChange={ (event) => setSerial(event.target.value) }
                             type="text"
                />
                <FormControl aria-label="wateringAmount" placeholder="مقدار آب (ml)"
                             onChange={ (event) => setAmount(event.target.value) }
                             type="text"
                />

            </InputGroup>
            <Button onClick={() => addNewWateringLog(deviceSerial, waterAmount)} variant="outline-secondary" id="button-addon2" className="add-device-btn set-form-btn">
                تنظیم
            </Button>
        </>
    );
}
export default Water;

import React, {useState} from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import {InputGroup, FormControl, Button, Modal} from 'react-bootstrap';
import "./humidity.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from "axios";
import useToken from "../../components/App/useToken";

const Humidity = () => {
    const token = useToken().token;

    const [deviceHumidity, setDeviceHumidity] = useState({
        device_serial: '',
        min_humidity: '',
        max_humidity: '',
    });
    const [isOpen, setIsOpen] = useState(false);
    const [success, setSuccess] = useState(false);

    const url_setHumidity = "http://usagi.carriot.ir:8000/device/set_humidity";
    const openModal = () => {
        setIsOpen(true);
        console.log(isOpen);
    }
    const closeModal = () => {
        setIsOpen(false);
    }


    const setHumidity = () =>{
        console.log(token, deviceHumidity)
        const deviceData = JSON.stringify({
            "device_serial" : deviceHumidity.device_serial,
            "min_humidity" : deviceHumidity.min_humidity,
            "max_humidity" : deviceHumidity.max_humidity
        });
        console.log(deviceData)
        axios.post(url_setHumidity, deviceData , {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
            .then((response) => {
                console.log(response)
                if (response.status == 200){
                    console.log("ajijam");
                    setSuccess(true);
                    openModal();
                } else {
                    closeModal();
                }
                //window.location.reload(true)

            });
    }

    return (
        <>
            <NavbarComp />
            <span className="gap"></span>
            
            <div className="page-title">
                تنظیم رطوبت
            </div>
            <InputGroup className="mb-3 humidity-input" size="lg">
                <FormControl aria-label="First name"
                             placeholder="سریال دستگاه"
                             onChange={(e) => setDeviceHumidity({...deviceHumidity, device_serial: e.target.value})}
                             type="text"
                />
            </InputGroup>
            <InputGroup dir="rtl" className="mb-3 humidity-input" size="lg">
                <FormControl aria-label="min humidity" placeholder="حداقل رطوبت"
                             onChange={(e) => setDeviceHumidity({...deviceHumidity, min_humidity: e.target.value})}
                             type="text" />
                <FormControl aria-label="max humidity" placeholder="حداکثر رطوبت"
                             onChange={(e) => setDeviceHumidity({...deviceHumidity, max_humidity: e.target.value})}
                             type="text" />


            </InputGroup>

            <Button variant="outline-secondary" id="button-addon2" className="add-device-btn set-form-btn"
                    onClick={() => setHumidity()}>
                تنظیم
            </Button>


            <Modal show={isOpen} onHide={closeModal}>
                    <Modal.Title className="success-modal" show={success}>تغییرات با موفقیت ذخیره شد.</Modal.Title>
                <Modal.Footer>
                    <Button  onClick={(e) => closeModal()} variant="outline-secondary" id="button-addon2" className="add-device-btn success-btn">
                        مرسی
                    </Button>
                </Modal.Footer>
            </Modal>

        </>
    );
}
export default Humidity;
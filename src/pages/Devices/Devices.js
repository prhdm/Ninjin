import React, { useState, useEffect } from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./devices.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from 'axios';
import useToken from "../../components/App/useToken";

const Devices = () => {
    const url_getAllDevice = 'http://usagi.carriot.ir/device/device_list';
    const url_deleteDevice = "http://usagi.carriot.ir/device/delete_device";
    const url_addDevice = "http://usagi.carriot.ir/device/create";
    const [devices, setdevices] = useState([])
    const [newDevice, setNewDevice] = useState({
        device_serial: "",
        device_name: "",
        phone: "",
        farm_id: 1

    })
    const token = useToken().token;


    const fetchData = () => {
        // console.log(token)
        axios.get(url_getAllDevice, {
            headers: {
              Authorization: 'Bearer ' + token
            }
           }).then(response => {
          setdevices(response.data.devices)
          console.log(response.data)
        })
      };
      

      useEffect(() => {
        fetchData()
      }, []);

      const deleteDevice = (serial) => {
        const deviceData = JSON.stringify({
            "deviceSerial" : serial
        });

        axios.post(url_deleteDevice, deviceData , {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
              } 
        })        
        .then((response) => {
            console.log(response)
        });
      }

      const addNewDevice = () =>{
        const deviceData = JSON.stringify({
            "device_serial" : "4933",
            "device_name" : "sara",
            "phone" : 3,
            "farm_id" : 1
        });

        axios.post(url_addDevice, deviceData , {
            headers: {
                'Authorization': `Bearer ${token}`
              } 
        })        
        .then((response) => {
            console.log(response)
        });

      }

    return (
        <>
            <NavbarComp />
            
            <span className="gap"></span>
            <div className="page-title">
                دستگاه‌ها
            </div>

            <InputGroup dir="rtl" className="mb-3 add-device">
                <InputGroup.Text id="basic-addon1">
                افرزودن دستگاه جدید
                </InputGroup.Text>
                
                <FormControl dir="rtl"
                placeholder="نام دستگاه"
                name="deviceName"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                onClick={() => deleteDevice(addNewDevice)}
                />
                <FormControl dir="rtl"
                placeholder="سریال دستگاه"
                name="deviceSerial"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                />
                <FormControl dir="rtl"
                placeholder="تلفن"
                name="phone"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                />
                <Button onClick={() => addNewDevice()} variant="outline-secondary" id="button-addon2" className="add-device-btn">
                افرزودن
                </Button>
            </InputGroup>

            <ListGroup as="ol" numbered>
                {devices.map(item => {
                return (
                    <>
                    <ListGroup.Item
                        as="li"
                        className="d-flex justify-content-between align-items-start"
                    >
                        <div className="ms-2 me-auto">
                            <div className="fw-bold">
                            {"نام دستگاه:\t"+item.deviceName}
                            </div>
                            <div>
                            {"سریال:\t" + item.deviceSerial}
                            </div>
                            {"تلفن:\t" + item.phone}
                        </div>
                        <Button onClick={() => deleteDevice(item.deviceSerial)} variant="outline-secondary" id="button-addon2" className="add-device-btn">
                            حذف
                        </Button>
                    </ListGroup.Item>
                    </>
                );
                })}
            </ListGroup>
        </>
    );
}

export default Devices;
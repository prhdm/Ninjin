import React, { useState, useEffect } from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./devices.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from 'axios';
import useToken from "../../components/App/useToken";

const Devices = () => {
    const url_getAllDevice = 'http://usagi.carriot.ir:8000/device/device_list';
    const url_deleteDevice = "http://usagi.carriot.ir:8000/device/delete_device";
    const url_addDevice = "http://usagi.carriot.ir:8000/device/create";
    const [devices, setdevices] = useState([])
    const [newDevice, setNewDevice] = useState({
        device_serial: "",
        device_name: "",
        phone: "",
        farm_id: 1

    })
    const token = useToken().token;


    const fetchData = () => {
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
                'Authorization': `Bearer ${token}`
              } 
        })        
        .then((response) => {
            console.log(response)
        });
      }

      const addNewDevice = () =>{
        const deviceData = JSON.stringify({
            "deviceSerial" : newDevice.device_serial,
            "deviceName" : newDevice.device_name,
            "phone" : newDevice.phone
        });
        console.log(deviceData)
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
                onChange={(e) => setNewDevice({...newDevice, device_name: e.target.value})}
                />
                <FormControl dir="rtl"
                placeholder="سریال دستگاه"
                name="deviceSerial"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                onChange={(e) => setNewDevice({...newDevice, device_serial: e.target.value})}
                />
                <FormControl dir="rtl"
                placeholder="تلفن"
                name="phone"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                onChange={(e) => setNewDevice({...newDevice, phone: e.target.value})}
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
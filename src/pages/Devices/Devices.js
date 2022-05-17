import React, { useState, useEffect } from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, InputGroup, FormControl, Button, Modal, Form} from 'react-bootstrap';
import "./devices.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from 'axios';
import useToken from "../../components/App/useToken";

const Devices = () => {
    const url_getAllDevice = 'http://usagi.carriot.ir:8000/device/device_list';
    const url_deleteDevice = "http://usagi.carriot.ir:8000/device/delete_device";
    const url_addDevice = "http://usagi.carriot.ir:8000/device/create";
    const url_editDevice = "http://usagi.carriot.ir:8000/device/edit-name"
    const [devices, setdevices] = useState([]);
    const [isOpen, setIsOpen] = useState(false);
    const [deviceNewName, setDeviceNewName] = useState ({
        device_serial:"",
        new_name:""
    });
    const [deviceSerial, setDeviceSerial] = useState("")

    const [newDevice, setNewDevice] = useState({
        device_serial: "",
        device_name: "",
        phone: "",
        farm_id: 1

    });
    const token = useToken().token;

    const openModal = (serial) => {
        setIsOpen(true);
        console.log(isOpen);
        setDeviceSerial(serial);
      }
      const closeModal = () => {
          setIsOpen(false);
        }

    const fetchData = () => {
        axios.get(url_getAllDevice, {
            headers: {
              Authorization: 'Bearer ' + token
            }
           }).then(response => {
          setdevices(response.data.devices)
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
            console.log(response);
            window.location.reload(true)
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
            window.location.reload(true)

        });
      }
///device/edit-name"
      const editDeviceName = () => {
        console.log(deviceSerial)
        console.log(deviceNewName.new_name)
        const deviceData = JSON.stringify({
            "device_serial" : deviceSerial,
            "new_name" : deviceNewName.new_name 
        });
        axios.post(url_editDevice, deviceData , {
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
                        <Button onClick={() => openModal(item.deviceSerial)} variant="outline-secondary" id="button-addon2" className="add-device-btn">
                            ویرایش
                            
                        </Button>
                        <span style={{"margin": "2px"}}></span>
                        <Button onClick={() => deleteDevice(item.deviceSerial)} variant="outline-secondary" id="button-addon2" className="add-device-btn">
                            حذف
                        </Button>
                        
                    </ListGroup.Item>

                    <Modal show={isOpen} onHide={closeModal}>
                        <Modal.Header closeButton>
                        <Modal.Title>ویرایش نام دستگاه</Modal.Title>
                        </Modal.Header>
                        <Form.Control onChange={(e) => setDeviceNewName({...deviceNewName, new_name: e.target.value})} type="name" className="inputNewName" placeholder="نام جدید را وارد کنید" />
                        <Modal.Footer>
                        <Button  onClick={(e) => editDeviceName()} variant="outline-secondary" id="button-addon2" className="add-device-btn">
                            ذخیره تغییرات
                        </Button>
                        </Modal.Footer>
                    </Modal>
                    
                    </>
                );
                })}
            </ListGroup>
        </>
    );
}

export default Devices;
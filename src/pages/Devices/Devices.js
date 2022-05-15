import React, { useState, useEffect } from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./devices.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from 'axios';
import useToken from "../../components/App/useToken";

//https://jsonplaceholder.typicode.com/users
const Devices = () => {
    const url = 'http://usagi.carriot.ir/device/device_list';
    const [users, setUsers] = useState([])
    const token = useToken().token;

    const fetchData = () => {
        console.log(token)
        axios.get(url, {
            headers: {
              Authorization: 'Bearer ' + token //the token is a variable which holds the token
            }
           }).then(response => {
          setUsers(response.data.devices)
          console.log(response.data)
        })
      }

      useEffect(() => {
        fetchData()
      }, [])

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
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                />
                <FormControl dir="rtl"
                placeholder="تلفن"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                />
                <Button variant="outline-secondary" id="button-addon2" className="add-device-btn">
                افرزودن
                </Button>
            </InputGroup>

            <ListGroup as="ol" numbered>
                {users.map(item => {
                return (
                    <>
                    <ListGroup.Item
                        as="li"
                        className="d-flex justify-content-between align-items-start"
                    >
                        <div className="ms-2 me-auto">
                            <div className="fw-bold">
                            {"name:\t"+item.deviceName}
                            </div>
                            <div>
                            {"serial:\t" + item.deviceSerial}
                            </div>
                            {"phone:\t" + item.phone}
                        </div>
                    </ListGroup.Item>
                    </>
                );
                })}
            </ListGroup>
        </>
    );
}

export default Devices;
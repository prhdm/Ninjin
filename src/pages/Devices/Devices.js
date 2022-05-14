import React, { useState, useEffect } from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, Badge, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./devices.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";

// usagi.carriot.ir/device/device_list -> device list
const Devices = () => {
    const [Users, fetchUsers] = useState([])

    const getData = () => {
        fetch('https://jsonplaceholder.typicode.com/users')
          .then((res) => res.json())
          .then((res) => {
            console.log(res)
            fetchUsers(res)
          })
      }
    
      useEffect(() => {
        getData()
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
                <Button variant="outline-secondary" id="button-addon2" className="add-device-btn">
                افرزودن
                </Button>
            </InputGroup>

            <ListGroup as="ol" numbered>
                {Users.map((item, i) => {
                return (
                    <>
                    <ListGroup.Item
                        as="li"
                        className="d-flex justify-content-between align-items-start"
                    >
                        <div className="ms-2 me-auto">
                            <div className="fw-bold">
                            {item.name}
                            </div>
                            <div>
                            {item.username}
                            </div>
                            {item.phone}
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
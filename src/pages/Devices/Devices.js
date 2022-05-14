import React from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, Badge, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./devices.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";

const Devices = () => {
    return (
        <>
            <NavbarComp />
            
            <span className="gap"></span>
            <div className="page-title">
                دستگاه‌ها
            </div>

            <InputGroup dir="rtl" className="mb-3 add-device">
                <FormControl dir="rtl"
                placeholder="نام دستگاه"
                aria-label="Recipient's username"
                aria-describedby="basic-addon2"
                />
                <Button variant="outline-secondary" id="button-addon2" className="add-device-btn">
                افرزودن
                </Button>
            </InputGroup>

            <ListGroup as="ol" numbered dir="rtl">
                <ListGroup.Item
                    as="li"
                    className="d-flex justify-content-between align-items-start"
                >
                    <div className="ms-2 me-auto">
                    <div className="fw-bold"></div>
                    Cras justo odio
                    </div>
                    <Badge bg="primary" pill>
                    0
                    </Badge>
                </ListGroup.Item>
                <ListGroup.Item
                    as="li"
                    className="d-flex justify-content-between align-items-start"
                >
                    <div className="ms-2 me-auto">
                    <div className="fw-bold">Subheading</div>
                    Cras justo odio
                    </div>
                    <Badge bg="primary" pill>
                    0
                    </Badge>
                </ListGroup.Item>
                <ListGroup.Item
                    as="li"
                    className="d-flex justify-content-between align-items-start"
                >
                    <div className="ms-2 me-auto">
                    <div className="fw-bold">Subheading</div>
                    Cras justo odio
                    </div>
                    <Badge bg="primary" pill>
                    0
                    </Badge>
                </ListGroup.Item>
            </ListGroup>
        </>
    );
}

export default Devices;
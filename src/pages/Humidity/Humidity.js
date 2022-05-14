import React from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, Badge, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./humidity.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";


const Humidity = () => {
    return (
        <>
            <NavbarComp />
            <span className="gap"></span>
            {/* <span className="gap"></span> */}
            <h3 className="humidity-title ml-3">تنظیم رطوبت</h3>
            <InputGroup className="mb-3 humidity-input" size="lg">
                {/* <InputGroup.Text>تنظیم رطوبت</InputGroup.Text> */}
                <FormControl aria-label="First name" placeholder="نام دستگاه" />
            </InputGroup>

            <InputGroup dir="rtl" className="mb-3 humidity-input" size="lg">
                {/* <InputGroup.Text>تنظیم رطوبت</InputGroup.Text> */}
                <FormControl aria-label="min humidity" placeholder="حداقل رطوبت" />
                <FormControl aria-label="max humidity" placeholder="حداکثر رطوبت" />

                    <Button variant="outline-secondary" id="button-addon2" className="add-device-btn">
                    تنظیم
                    </Button>
            </InputGroup>
        </>
    );
}
export default Humidity;
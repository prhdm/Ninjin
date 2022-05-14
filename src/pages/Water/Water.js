import React from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { InputGroup, FormControl, Button} from 'react-bootstrap';
import "bootstrap/dist/css/bootstrap.rtl.min.css";


const Water = () => {
    return (
        <>
            <NavbarComp />
            <span className="gap"></span>
            
            <div className="page-title">
                تنظیم مقدار آبدهی
            </div>

            <InputGroup dir="rtl" className="mb-3 humidity-input" size="lg">
                <FormControl aria-label="max humidity" placeholder="نام دستگاه" />
                <FormControl aria-label="min humidity" placeholder="مقدار آب" />

                    <Button variant="outline-secondary" id="button-addon2" className="add-device-btn">
                    تنظیم
                    </Button>
            </InputGroup>
        </>
    );
}
export default Water;
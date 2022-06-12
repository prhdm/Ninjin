import React from "react";
import { ProSidebar, Menu, MenuItem, SubMenu } from 'react-pro-sidebar';
import 'react-pro-sidebar/dist/css/styles.css';
import AvatarEditor from 'react-avatar-editor';
import profile from "../components/Navbar/img/usagi pofile.jpg";
import './slidebar.css';

const Slidebar = (props) => {
    const handleLogout = () => {
        localStorage.clear();
        window.location.reload();
    }
    return (
        <>
            <ProSidebar>
            <Menu iconShape="square" className="menu">
                                <AvatarEditor className="avatar"
                    image={profile}
                    width={76}
                    height={76}
                    border={0}
                    color={[255, 255, 255, 0.6]} // RGBA
                    scale={1.2}
                    rotate={0}
                />
                <MenuItem className="item">گزارش</MenuItem>
                <MenuItem>لیست دستگاه ها</MenuItem>
                <MenuItem>گزارش رطوبت</MenuItem>
                <MenuItem>اخطارها</MenuItem>
                <MenuItem>تنظیمات</MenuItem>
                <MenuItem>
                    <button className="btn logout" onClick={handleLogout}>خروج</button>
                </MenuItem>
                <SubMenu title="Components">
                <MenuItem>Component 1</MenuItem>
                <MenuItem>Component 2</MenuItem>
                </SubMenu>
            </Menu>
            </ProSidebar>;

        </>
    );
}

export default Slidebar;
import React from "react";
import profile from "./img/usagi pofile.jpg";
import AvatarEditor from 'react-avatar-editor';
import "./avatar.css";

const UserAvatar = () => {
    return (
        <>
                 <AvatarEditor className="avatar"
                    image={profile}
                    width={76}
                    height={76}
                    border={0}
                    color={[255, 255, 255, 0.6]} // RGBA
                    scale={1.2}
                    rotate={0}
                />
        </>
    );
}

export default UserAvatar;
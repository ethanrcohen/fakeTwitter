import React from 'react';
import UserDisplay from './UserDisplay';
import { Avatar } from '@material-ui/core';


type PostProps = {
    content: string;
    pictureFile: string;
};

const Post = (props: PostProps) => {
    return (
        // TODO: Make this look not awful
        <div className="flex">
            <Avatar alt="Avatar" src={props.pictureFile} />
            <div>
                <UserDisplay name="React JS" handle="reactjs" ></UserDisplay>
                <p>{props.content}</p>
            </div>
        </div>
    );
};

export default Post;
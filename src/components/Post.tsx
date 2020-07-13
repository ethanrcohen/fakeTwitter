import React from 'react';
import UserDisplay from './UserDisplay';
import logo from '../images/logo.svg';

type PostProps = {
    content: string;
};

class Post extends React.Component<PostProps> {
    content: string;

    constructor(props: PostProps) {
        super(props);
        this.content = props.content;
    }

    render() {
        return (
            // TODO: Make this look not awful
            <div>
                <UserDisplay name="React JS" handle="reactjs" pictureFile={logo} ></UserDisplay>
                <p>{this.content}</p>
            </div>
        );
    }
}

export default Post;
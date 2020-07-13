import React from 'react';
import { Avatar } from '@material-ui/core';

type UserDisplayProps = {
    name: string,
    handle:string,
    pictureFile: string
};

class UserDisplay extends React.Component<UserDisplayProps> {
    name: string;
    handle: string;
    pictureFile: string;

    constructor(props: UserDisplayProps) {
        super(props);
        this.name = props.name;
        this.handle = props.handle;
        // TODO: overwrite with a default image if not provided
        this.pictureFile = props.pictureFile;
    }

    render() {
        return (
            <div>
                <Avatar alt="Avatar" src={this.pictureFile}/>
                <p><b>{this.name}</b> @{this.handle}</p>
            </div>
        );
    }
}

export default UserDisplay;
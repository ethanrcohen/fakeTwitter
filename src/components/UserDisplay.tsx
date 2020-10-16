import React from 'react';

type UserDisplayProps = {
    name: string,
    handle: string,
};

const UserDisplay = (props: UserDisplayProps) => {
    return (
        <p><b>{props.name}</b> @{props.handle}</p>
    );
};

export default UserDisplay;
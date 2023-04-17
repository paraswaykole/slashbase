export enum ProfileImageSize {
    TINY = "TINY",
    SMALL = "SMALL",
    MEDIUM = "MEDIUM",
    LARGE = "LARGE"
}

type ProfileImagePropType = {
    imageUrl: string,
    classes?: string[],
    size?: ProfileImageSize
}

const ProfileImage = ({ imageUrl, classes, size }: ProfileImagePropType) => {


    if (!size) {
        size = ProfileImageSize.SMALL
    }

    let width, height
    if (size === ProfileImageSize.LARGE) {
        width = 100
        height = 100
    } else if (size === ProfileImageSize.MEDIUM) {
        width = 70
        height = 70
    } else if (size === ProfileImageSize.SMALL) {
        width = 40
        height = 40
    } else {
        width = 25
        height = 25
    }

    if (imageUrl === '')
        return (
            <div style={{ width, height }} className={["profileImageDefault", ...(classes ?? [])].join(' ')}>
                <i className={"fas fa-user " + (size !== ProfileImageSize.SMALL ? 'fa-2x' : '')}></i>
            </div>
        )

    return (
        <img className={["profileImage", ...(classes ?? [])].join(' ')} src={imageUrl} width={width} height={height} />
    )
}


export default ProfileImage


import Button from "../ui/Button"

type AddModal = {
    message: string,
    onConfirm: () => void
    onClose: () => void
}

const ConfirmModal = ({ message, onConfirm, onClose }: AddModal) => {

    return (
        <div className="modal is-active">
            <div className="modal-background" onClick={onClose}></div>
            <div className="modal-content" style={{ width: 'initial' }}>
                <div className="box">
                    <h2>{message}</h2><br />
                    <Button className="is-small is-primary" onClick={onConfirm}>Confirm</Button>&nbsp;
                    <Button className="is-small" onClick={onClose}>Cancel</Button>
                </div>
            </div>
        </div >
    )
}

export default ConfirmModal

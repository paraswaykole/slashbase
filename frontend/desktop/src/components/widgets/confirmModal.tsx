
import React from 'react'

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
                    <button className="button is-small is-primary" onClick={onConfirm}>Confirm</button>&nbsp;
                    <button className="button is-small " onClick={onClose}>Cancel</button>
                </div>
            </div>
        </div >
    )
}

export default ConfirmModal

import { useLocation, useNavigate } from 'react-router-dom'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import styles from './footer.module.scss'
import Constants from '../../constants'
import { useEffect } from 'react'
import { checkConnection, selectDBConnection, selectIsDBConnected, getDBDataModels, resetDBDataModels } from '../../redux/dbConnectionSlice'


type FooterPropType = {}

const Footer = (_: FooterPropType) => {

    const navigate = useNavigate()
    const location = useLocation()
    const dispatch = useAppDispatch()

    const showStatus = location.pathname.startsWith("/db")

    const dbConnection = useAppSelector(selectDBConnection)
    const isDBConnected = useAppSelector(selectIsDBConnected)

    useEffect(() => {
        if (showStatus && dbConnection) {
            dispatch(checkConnection())
        }
    }, [showStatus, dbConnection])

    const openSupport = () => {
        navigate(Constants.APP_PATHS.SETTINGS_SUPPORT.path)
    }

    const refreshDataModels = () => {
        dispatch(resetDBDataModels())
        dispatch(getDBDataModels({ dbConnId: dbConnection!.id }))
    }

    return (
        <footer className={styles.footer}>
            <div>
                {showStatus && isDBConnected !== undefined &&
                    (<button className={styles.button + " is-small"}>
                        <span className="icon is-small">
                            {!isDBConnected && <i className="far fa-circle" />}
                            {isDBConnected && <i className="fas fa-circle" />}
                        </span>
                        <span>{(isDBConnected !== undefined && isDBConnected) ? "connected" : "not connected"}</span>
                    </button>)
                }
                {showStatus && isDBConnected === true &&
                    (<button className={styles.button + " is-small"} onClick={refreshDataModels}>
                        <span className="icon is-small">
                            <i className="fas fa-sync" />
                        </span>
                        <span>refresh data models</span>
                    </button>)
                }
            </div>
            <div>
                <button className={styles.button + " is-small"} onClick={openSupport}>
                    <span className="icon is-small">
                        <i className="far fa-circle-question" />
                    </span>
                    <span>Help & Feedback</span>
                </button>
            </div>
        </footer>
    )
}


export default Footer
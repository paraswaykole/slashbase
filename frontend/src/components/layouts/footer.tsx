import { selectIsConnected } from '../../redux/apiSlice'
import { useAppSelector } from '../../redux/hooks'
import styles from './footer.module.scss'


type FooterPropType = {}

const Footer = (_: FooterPropType) => {

    const isConnected = useAppSelector(selectIsConnected)

    return (
        <footer className={styles.footer}>
            {!isConnected && <span>🔴 not connected</span>}
            {isConnected && <span>🟢 connected</span>}
        </footer>
    )
}


export default Footer
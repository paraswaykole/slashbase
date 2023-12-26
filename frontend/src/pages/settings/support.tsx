import React, { FunctionComponent } from 'react'
import Constants from '../../constants'
import utils from '../../lib/utils'
import Button from '../../components/ui/Button'

const SupportPage: FunctionComponent<{}> = () => {

    return (
        <React.Fragment>

            <h1>Support</h1>

            <p>We are more than happy to help you.</p>
            <br />
            <h4>Have questions/ideas or need help?</h4>
            <p>If you are have any questions or ideas,
                please join our Discord server and share. We are open to discussing anything your questions or ideas.</p>
            <a onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.DISCORD_COMMUNITY) }}>
                <Button
                    text='Join Discord Community'
                    className="is-secondary"
                    icon={<i className={"fab fa-discord"} />}
                />
            </a>
            <br /><br />
            <h4>Facing bugs or errors?</h4>
            <p>If you are getting any errors or bugs, join our Discord server and ask for help.
                We will try to fix the errors bugs in the next version release.</p>
            <a onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.REPORT_BUGS) }}>
                <Button
                    text='File a Github Issue'
                    className="is-secondary"
                    icon={<i className={"fab fa-github"} />}
                />
            </a>
            <br /><br />
            <h4>Want to check what&apos;s new?</h4>
            <p>Vist our website to check new releases.</p>
            <a onClick={() => { utils.openInBrowser(Constants.EXTERNAL_PATHS.OFFICIAL_WEBSITE) }}>
                <Button
                    text='Visit Slashbase.com'
                    className="is-secondary"
                    icon={<i className='fas fa-globe' />}
                />
            </a>
        </React.Fragment>
    )
}

export default SupportPage

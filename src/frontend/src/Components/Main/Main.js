import React from 'react';
import PageTypes from '../../Constants/PageTypes/PageTypes';
import MainPageContent from './Content/MainPageContent/MainPageContent';
import SignOutButton from './Components/SignOutButton/SignOutButton';
import UpdateName from './Components/UpdateName/UpdateName';
import UpdateAvatar from './Components/UpdateAvatar/UpdateAvatar';
import Analysis from './Content/MainPageContent/Analysis';
import AllTransaction from './Content/MainPageContent/AllTransaction';

const Main = ({ page, setPage, setAuthToken, setUser, user , authToken}) => {
    let content = <></>
    let contentPage = true;
    switch (page) {
        case PageTypes.signedInMain:
            content = <MainPageContent user={user} setAuthToken={setAuthToken} authToken={authToken} setPage={setPage} />;
            break;
        case PageTypes.signedInUpdateName:
            content = <UpdateName user={user} setUser={setUser} />;
            break;
        case PageTypes.signedInUpdateAvatar:
            content = <UpdateAvatar user={user} setUser={setUser} />;
            break;
        case PageTypes.analysis:
            content = <Analysis user={user}/>;
            break;
        case PageTypes.allTransactions:
            content = <AllTransaction user={user}/>;
            break;
        default:
            content = <>Error, invalid path reached</>;
            contentPage = false;
            break;
    }
    return <>
        <ul>
            <li>{contentPage && <button onClick={(e) => setPage(e, PageTypes.signedInMain)}>Back to main</button>}</li>
            <li><SignOutButton setUser={setUser} setAuthToken={setAuthToken} /></li>
        </ul>
        {content}
    </>
}

export default Main;
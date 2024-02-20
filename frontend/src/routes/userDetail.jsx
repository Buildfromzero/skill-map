import UserList from '../components/userList';
import Navbar from '../components/navbar';
import UserDetail from '../components/userDetail';

function UserPage() {
    return (

        <div className="container-fluid">

            <div className="row">
                <div className="col gap-3">
                    <Navbar />
                </div>
            </div>

            <br />

            <nav aria-label="breadcrumb">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><a href="/">Home</a></li>
                    <li class="breadcrumb-item active" aria-current="users">Users</li>
                </ol>
            </nav>

            <br />

            <div className="row align-items-center">
                <div className="col-4">
                    <UserDetail />
                </div>
            </div>


            <br />
        </div>
    );
}

export default UserPage;

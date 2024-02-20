import Navbar from '../components/navbar';
import SkillList from '../components/skillList';
import CreateSkill from '../components/skillCreate';

function SkillPage() {
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
                    <li class="breadcrumb-item active" aria-current="users">Skills</li>
                </ol>
            </nav>

            <br />

            <div className="row align-items-center">
                <div className="col-4">
                    <CreateSkill />
                </div>
            </div>

            <br />

            <div className="row">
                <div className="col">
                    <SkillList />
                </div>
            </div>

            <br />
        </div>
    );
}

export default SkillPage;

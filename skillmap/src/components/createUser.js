export default function CreateUser() {
    function onSubmission(event){
        console.log(event.target.elements.firstName.value)
        console.log(event.target.elements.lastName.value)
    }
    return (
      <div>
        <h3>Create User</h3>
            <form onSubmit={onSubmission}>
                <div class="mb-3">
                    <label class="form-label">First Name</label>
                    <input type="text" class="form-control" id="firstName"/>
                </div>
                <div class="mb-3">
                    <label class="form-label">Last Name</label>
                    <input type="text" class="form-control" id="lastName"/>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </div>
    );
  }
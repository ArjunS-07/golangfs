function CarView(){
    return(
        <>
        <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
    <div className="container-fluid">
      <a className="navbar-brand" href="cars_list.html">Parking Management System</a>
      <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>
      
      <div className="collapse navbar-collapse" id="navbarSupportedContent">
        <ul className="navbar-nav me-auto mb-2 mb-lg-0">
          <li className="nav-item">
            <a className="nav-link " aria-current="page" href="cars_list.html">Cars List</a>
          </li>
          <li className="nav-item">
            <a className="nav-link active" href="cars_create.html">Add Cars</a>
          </li>

        </ul>

      </div>
    </div>
  </nav>
  <h3><a href="cars_list.html" className="btn btn-light">Go Back</a>View Car</h3>
  <div className="container">
    <div className="form-group mb-3">
      <label for="number" className="form-label">Car Number:</label>
      <div className="form-control" id="number">???</div>
    </div>
    <div className="form-group mb-3">
      <label for="model" className="form-label">Car Model:</label>
      <div className="form-control" id="model">???</div>
    </div>
    <div className="form-group mb-3">
      <label for="type" className="form-label">Car Type(SUV/ CUV/ Sedan):</label>
      <div className="form-control" id="type">???</div>
    </div>
  </div>
        </>
    )
}
export default CarView;
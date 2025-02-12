function CarsList(){
  return(
    <>
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
        <div className="container-fluid">
          <a className="navbar-brand" href="cars_list.html">Parking Management System</a>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarSupportedContent">
            <ul className="navbar-nav me-auto mb-2 mb-lg-0">
              <li className="nav-item">
                <a className="nav-link active" aria-current="page" href="cars_list.html">Cars List</a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="cars_create.html">Add cars</a>
              </li>
              
            </ul>
            
          </div>
        </div>
    </nav>
    <h3>Car List</h3>
    <div className="container">
        <table className="table table-success table-striped">
        <thead>
            <tr>
              <th scope="col">id</th>
              <th scope="col">number</th>
              <th scope="col">model</th>
              <th scope="col">class</th>
              <th scope="col"></th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <th scope="row">car1</th>
              <td>KA09asc1231</td>
              <td>Zen</td>
              <td>smallcar</td>
              <td><a className="btn btn-success" href="car_view.html?id=car1">view</a></td>
            </tr>
            <tr>
              <th scope="row">car2</th>
              <td>KA09asc3211</td>
              <td>Beatle</td>
              <td>hatchback</td>
              <td><a className="btn btn-success" href="car_view.html?car2">view</a></td>
            </tr>
        </tbody>
        </table>
    </div>
    </>
  )
}
export default CarsList;
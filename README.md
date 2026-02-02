This project is a cinema booking backend API built with GO(Golang),fiber, Gorm and JWT authentication. It allows user to browse movie and showtimes,reserve seats, manage their bookings, and supports role-based access control for admins.
Features:
-Authentication and Authoriztion:JWT-based authentication_ Protected routs using middleware_Role-based access control
-Movies and Showtimes: Get all movies_Get all showtimes_ Get showtime by ID_ Automatically ignore expired showtimes
-Seats: Seats belong to a hall_ Seats are fetched per showtime _ Each seat shows its availability status 
-Booking system: Users can reserve one or multiple seats for a specific showtime _ Cancelled bookins free the seats again _ Users can only cancel their own bookings _ Admins can view all bookings
--Booking logic:
-A seat is considered booked only if: (status="reserved")
-if a booking status is: (status="cancel") -> the seats become available again.
--Booking flow: 
1_User logs in and receives JWT token.
2_User fetches showtimes.
3_User fetches seats by showtime.
4_User reserves one or more seats.
5_User can view their own bookings.
6_User can cancel a booking.
--API endpoints:
-Auth: POST /user/register
       POST /user/login
       POST /make-admin/:id
-Showtimes:Get /showtime
           Get /showtime/id
           POST /admin/showtime
           PUT /admin/showtime/{id}
           DEL /admin/showtime/{id}
-Seats: GET /seat
        GET /seat/hall_id
        GET /showtime/1/seats
        POST /admin/seat
        PUT /admin/seat/seat_id
        DELETE /admin/seat/seat_id
        PATCH /admin/seat/23/status
-Booking: POST /book-seats
          GET /my-bookings
          PATCH /booking/2/cancel
--How to run a project: _git clone https://github.com/your-username/your-repo-name.git
                        _cd your-repo-name
                        _go mod tidy
                        _go run main.go
                        -Make sure to configure your database connection and JWT secret in the environment variables before running the project.
--Authentication and Authoriztion:-Authentication is handled using JWT tokens.
                                  -Protected routs require an Authorization header :   Authorization: Bearer <your_token>
--Role-based access control:_User: can browse showtime, reserve seats, view and cancel own bookings.
                            _Admin: can manage movies, showtimes, seats, and view all bookings.
                            _Super Admin:can promote users to admin.
                        

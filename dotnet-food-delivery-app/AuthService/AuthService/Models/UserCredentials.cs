namespace AuthService.Models
{
    public enum UserRole 
    {
        ADMIN, 
        CUSTOMER
    }
    public class UserCredentials
    {
        public Guid Id { get; set; }

        public UserRole UserRole { get; set; }

        public int Email { get; set; }

        public int Password { get; set; }
    }
}

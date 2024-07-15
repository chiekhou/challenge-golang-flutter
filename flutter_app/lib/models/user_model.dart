class User {
  int id;
  String firstName;
  String lastName;
  String username;
  String? photo;
  String email;
  String? password;
  String address;
  int roleId;

  User({
    required this.id,
    this.photo,
    required this.firstName,
    required this.lastName,
    required this.username,
    required this.email,
    this.password,
    required this.address,
    required this.roleId,
  });

  @override
  String toString() {
    return 'User{id: $id, first_name: $firstName, last_name: $lastName, '
        'username: $username, email: $email, password: $password, photo: $photo,'
        'address: $address, role_id: $roleId}';
  }

  User.fromJson(Map<String, dynamic> json)
      : id = json['id'] ?? 0,
        firstName = json['first_name'] ?? '',
        lastName = json['last_name'] ?? '',
        username = json['username'] ?? '',
        password = json['password'] ?? '',
        photo = json['photo'] ?? '',
        email = json['email'] ?? '',
        address = json['address'] ?? '',
        roleId = json['role_id'] ?? 0;

  Map<String, dynamic> toJson() {
    Map<String, dynamic> value = {
      'id': id,
      'first_name': firstName,
      'last_name': lastName,
      'username': username,
      'password': password,
      'photo': photo,
      'email': email,
      'address': address,
      'role_id': roleId
    };
    value['id'] = id;
    return value;
  }
}

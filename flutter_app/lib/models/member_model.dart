class Member{
  int id;
  String first_name;
  String last_name;
  String username;
  String? photo;
  String email;
  String? password;
  String address;
  Member({
    required this.id,
    this.photo,
    required this.first_name,
    required this.last_name,
    required this.username,
    required this.email,
    this.password,
    required this.address,
});

  @override
  String toString(){
    return 'User{id: $id, first_name: $first_name, last_name: $last_name, '
        'username: $username, email: $email, password: $password, photo: $photo,'
        'address: $address}';
  }

  Member.fromJson(Map<String, dynamic> json) :
        id = json['id'] ?? 0,
        first_name = json['first_name'] ?? '',
        last_name = json['last_name'] ?? '',
        username = json['username'] ?? '',
        password = json['password'] ?? '',
        photo = json['photo'] ?? '',
        email = json['email'] ?? '',
        address = json['address'] ?? '';

  Map<String, dynamic> toJson(){
    Map<String, dynamic> value = {
      'id': id,
      'first_name': first_name,
      'last_name': last_name,
      'username': username,
      'password': password,
      'photo': photo,
      'email': email,
      'address': address
    };
    value['id'] = id;
    return value;
  }
}
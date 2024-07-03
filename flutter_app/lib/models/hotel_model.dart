enum HotelStatus { ongoing, done }

class Hotel {
  int id;
  String name;
  String image;
  String destination;
  double price;
  HotelStatus status;
  LocationHotel? location;
  Hotel({
    required this.name,
    required this.destination,
    required this.image,
    required this.price,
    required this.id,
    this.location,
    this.status = HotelStatus.ongoing,
  });


  String toString() {
    return 'Hotel{id: $id, name: $name,destination: $destination, price: $price}';
  }

  Hotel.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        name = json['name'],
        destination = json['destination'],
        image = json['image'],
        price = json['price'].toDouble(),
        location = LocationHotel(
          address: json['address'],
          latitude: json['latitude'],
          longitude: json['longitude'],
  ),
  status =
  json['status'] == 0 ? HotelStatus.ongoing : HotelStatus.done;

  Map<String, dynamic> toJson() {
    Map<String, dynamic> value = {
      'id': id,
      'name': name,
      'image': image,
      'price': price,
      'address': location?.address,
      'longitude': location?.longitude,
      'latitude': location?.latitude,
      'status': status == HotelStatus.ongoing ? 0 : 1
    };
    value['id'] = id;
    return value;
  }
}

class LocationHotel {
  String? address;
  double? longitude;
  double? latitude;
  LocationHotel({
    this.address,
    this.longitude,
    this.latitude,
  });
}



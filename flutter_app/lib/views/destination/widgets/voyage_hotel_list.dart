import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import '../../../models/activity_model.dart';

class VoyageHotelList extends StatelessWidget {
  final List<Hotel> hotels;
  final Function deleteVoyageHotel;

  const VoyageHotelList(
      {super.key,
        required this.hotels,
        required this.deleteVoyageHotel});

  @override
  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      physics: NeverScrollableScrollPhysics(),
      shrinkWrap: true,
      itemBuilder: (context, index) {
        var hotel = hotels[index];
        return Card(
          child: ListTile(
            leading: CircleAvatar(
              backgroundImage: NetworkImage(hotel.image),
            ),
            title: Text(hotel.name),
            trailing: IconButton(
              icon: const Icon(
                Icons.delete,
                color: Colors.red,
              ),
              onPressed: () {
                deleteVoyageHotel(hotels[index]);
                ScaffoldMessenger.of(context)
                  ..removeCurrentSnackBar()
                  ..showSnackBar(
                    const SnackBar(
                      content: Text('Hôtel supprimé'),
                      backgroundColor: Colors.red,
                      duration: Duration(seconds: 1),
                    ),
                  );
              },
            ),
          ),
        );
      },
      itemCount: hotels.length,
    );
  }
}
import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

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
      physics: const NeverScrollableScrollPhysics(),
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
                    SnackBar(
                      content: Text(AppLocalizations.of(context)!.delete_hotel),
                      backgroundColor: Colors.red,
                      duration: const Duration(seconds: 1),
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
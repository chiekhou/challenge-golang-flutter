import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import 'package:flutter_app/providers/voyage_provider.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:provider/provider.dart';
import 'package:url_launcher/url_launcher.dart';

class GoogleMapViewHotel extends StatefulWidget {
  static const String routeName = '/google-map-hotel';

  const GoogleMapViewHotel({super.key});

  @override
  State<GoogleMapViewHotel> createState() => _GoogleMapViewHotelState();
}

class _GoogleMapViewHotelState extends State<GoogleMapViewHotel> {
  final bool _isLoaded = false;
  late GoogleMapController _controller;
  late Hotel _hotel;

  @override
  void didChangeDependencies() {
    if (!_isLoaded) {
      var arguments =
      ModalRoute.of(context)!.settings.arguments as Map<String, int?>;
      _hotel =
          Provider.of<VoyageProvider>(context, listen: false).getHotelByIds(
            hotelId: arguments['hotelId'] as int,
            voyageId: arguments['voyageId'] as int,
          );
      if (_hotel.location == null) {
        Navigator.pop(context, null);
      }
    }
    super.didChangeDependencies();
  }

  get _hotelLatLng {
    return LatLng(
      _hotel.location!.latitude!,
      _hotel.location!.longitude!,
    );
  }

  get _initialCameraPosition {
    return CameraPosition(
      target: _hotelLatLng,
      zoom: 16.0,
    );
  }

  Future<void> _openUrl() async {
    Uri url = Uri.parse('google.navigation:q=${_hotel.location!.address}');
    if (await canLaunchUrl(url)) {
      await launchUrl(url);
    } else {
      throw 'cannot launch url';
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(_hotel.name),
      ),
      body: GoogleMap(
        initialCameraPosition: _initialCameraPosition,
        mapType: MapType.normal,
        onMapCreated: (controller) => _controller = controller,
        markers: {
          Marker(
            markerId: const MarkerId('123'),
            flat: true,
            position: _hotelLatLng,
          )
        },
      ),
      floatingActionButton: FloatingActionButton.extended(
        icon: const Icon(Icons.directions_car),
        onPressed: _openUrl,
        label: const Text('Go'),
      ),
    );
  }
}

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../../apis/google_api.dart';
import '../../../models/activity_model.dart';
import '../../../providers/destination_provider.dart';
import 'activity_form_autocomplete.dart';
import 'activity_form_image_picker.dart';
import 'package:location/location.dart';

class ActivityForm extends StatefulWidget {
  final String destinationName;

  const ActivityForm({super.key, required this.destinationName});

  @override
  State<ActivityForm> createState() => _ActivityFormState();
}

class _ActivityFormState extends State<ActivityForm> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  late FocusNode _priceFocusNode;
  late FocusNode _urlFocusNode;
  late FocusNode _addressFocusNode;
  late Activity _newActivity;
  late String? _nameInputAsync;
  final TextEditingController _urlController = TextEditingController();
  final TextEditingController _addressController = TextEditingController();
  bool _isLoading = false;
  FormState get form {
    return _formKey.currentState!;
  }

  @override
  void initState() {
    _newActivity = Activity(
      id: 0,
      destination: widget.destinationName,
      name: '',
      price: 0,
      image: '',
      location: LocationActivity(
        address: null,
        longitude: null,
        latitude: null,
      ),
      status: ActivityStatus.ongoing,
    );
    _priceFocusNode = FocusNode();
    _urlFocusNode = FocusNode();
    _addressFocusNode = FocusNode();
    _addressFocusNode.addListener(() async {
      if (_addressFocusNode.hasFocus) {
        var location = await showInputAutocomplete(context);
        _addressFocusNode.nextFocus();
        if (location != null) {
          _newActivity.location = location;
          setState(() {
            if (location != null) {
              _addressController.text = location.address!;
            }
          });
          _urlFocusNode.requestFocus();
        }
      } else {
        print('no focus');
      }
    });
    super.initState();
  }

  void updateUrlField(String url) {
    setState(() {
      _urlController.text = url;
    });
  }

  void _getCurrentLocation() async {
    try {
      Location location = Location();
      LocationData userLocation;
      bool serviceEnabled;
      PermissionStatus permissionGranted;

      serviceEnabled = await location.serviceEnabled();
      if (!serviceEnabled) {
        serviceEnabled = await location.requestService();
        if (!serviceEnabled) {
          return;
        }
      }
      permissionGranted = await location.hasPermission();
      if (permissionGranted == PermissionStatus.denied) {
        permissionGranted = await location.requestPermission();
        if (permissionGranted != PermissionStatus.granted) {
          return;
        }
      }
      userLocation = await location.getLocation();

      String address = await getAddressFromLatLng(
        lat: userLocation.latitude!,
        lng: userLocation.longitude!,
      );
      _newActivity.location = LocationActivity(
        address: address,
        latitude: userLocation.latitude!,
        longitude: userLocation.longitude!,
      );
      setState(() {
        _addressController.text = address;
      });
    } catch (e) {
      rethrow;
    }
  }

  @override
  void dispose() {
    _priceFocusNode.dispose();
    _urlFocusNode.dispose();
    _addressFocusNode.dispose();
    _urlController.dispose();
    _addressController.dispose();
    super.dispose();
  }

  Future<void> submitForm() async {
    try {
      DestinationProvider destinationProvider =
          Provider.of<DestinationProvider>(
        context,
        listen: false,
      );
      _formKey.currentState!.save();
      setState(() => _isLoading = true);
      _nameInputAsync = await destinationProvider.verifyIfActivityNameIsUnique(
        widget.destinationName,
        _newActivity.name,
      );
      if (form.validate()) {
        await destinationProvider.addActivityToDestination(_newActivity);
        if (mounted) Navigator.pop(context);
      } else {
        setState(() => _isLoading = false);
      }
    } catch (e) {
      setState(() => _isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(15),
      child: Form(
        key: _formKey,
        child: Column(
          children: <Widget>[
            TextFormField(
              autofocus: true,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Remplissez le nom';
                } else if (_nameInputAsync != null) {
                  return _nameInputAsync;
                }
                return null;
              },
              textInputAction: TextInputAction.next,
              decoration: const InputDecoration(
                labelText: 'Nom',
              ),
              onSaved: (value) => _newActivity.name = value!,
              onFieldSubmitted: (_) =>
                  FocusScope.of(context).requestFocus(_priceFocusNode),
            ),
            const SizedBox(
              height: 30,
            ),
            TextFormField(
              keyboardType: TextInputType.number,
              textInputAction: TextInputAction.next,
              focusNode: _priceFocusNode,
              decoration: const InputDecoration(
                hintText: 'Prix',
              ),
              onFieldSubmitted: (_) =>
                  FocusScope.of(context).requestFocus(_urlFocusNode),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Remplissez le Prix';
                }
                return null;
              },
              onSaved: (value) => _newActivity.price = double.parse(value!),
            ),
            const SizedBox(
              height: 30,
            ),
            TextFormField(
              focusNode: _addressFocusNode,
              controller: _addressController,
              decoration: const InputDecoration(
                hintText: 'Adresse',
              ),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Remplissez une adresse valide';
                }
                return null;
              },
              onSaved: (value) => _newActivity.location!.address = value!,
            ),
            const SizedBox(
              height: 10,
            ),
            TextButton.icon(
              icon: const Icon(Icons.gps_fixed),
              label: const Text('Utilisez ma position actuelle'),
              onPressed: _getCurrentLocation,
            ),
            const SizedBox(
              height: 30,
            ),
            TextFormField(
              keyboardType: TextInputType.url,
              focusNode: _urlFocusNode,
              controller: _urlController,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Remplissez l\'url';
                }
                return null;
              },
              decoration: const InputDecoration(
                hintText: 'Url image',
              ),
              onSaved: (value) => _newActivity.image = value!,
            ),
            const SizedBox(
              height: 30,
            ),
            ActivityFormImagePicker(updateUrl: updateUrlField),
            const SizedBox(
              height: 30,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: <Widget>[
                TextButton(
                  child: const Text('annuler'),
                  onPressed: () => Navigator.pop(context),
                ),
                ElevatedButton(
                  onPressed: _isLoading ? null : submitForm,
                  child: const Text('sauvegarder'),
                ),
              ],
            )
          ],
        ),
      ),
    );
  }
}

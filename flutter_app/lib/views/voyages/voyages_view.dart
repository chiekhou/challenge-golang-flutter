import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../providers/voyage_provider.dart';
import '../../widgets/app_loader.dart';
import 'widgets/voyage_list.dart';
import '../../widgets/app_drawer.dart';

class VoyagesView extends StatelessWidget {
  static const String routeName = '/voyages';

  const VoyagesView({super.key});

  @override
  Widget build(BuildContext context) {
    VoyageProvider voyageProvider = Provider.of<VoyageProvider>(context);
    return DefaultTabController(
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Mes voyages'),
          bottom: const TabBar(
            tabs: <Widget>[
              Tab(
                text: 'A venir',
              ),
              Tab(
                text: 'Passés',
              ),
            ],
          ),
        ),
        drawer: const AppDrawer(),
        body: voyageProvider.isLoading == false
            ? voyageProvider.voyages.isNotEmpty
                ? TabBarView(
                    children: <Widget>[
                      VoyageList(
                        voyages: voyageProvider.voyages
                            .where((voyage) =>
                                DateTime.now().isBefore(voyage.date!))
                            .toList(),
                      ),
                      VoyageList(
                        voyages: voyageProvider.voyages
                            .where((voyage) =>
                                DateTime.now().isAfter(voyage.date!))
                            .toList(),
                      ),
                    ],
                  )
                : Container(
                    alignment: Alignment.center,
                    child: const Text('Aucun voyage pour le moment'),
                  )
            : const AppLoader(),
      ),
    );
  }
}

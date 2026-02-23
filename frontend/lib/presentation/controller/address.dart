import 'package:flutter/cupertino.dart';
import 'package:frontend/domain/usecase/address.dart';

import '../../domain/entity/address.dart';

class AddressController{
  final GetAddressUseCase getAddressUseCase;

  AddressController({required this.getAddressUseCase});

  final ValueNotifier<bool> isLoading = ValueNotifier<bool>(false);
  final ValueNotifier<String> isError = ValueNotifier<String>("");
  final ValueNotifier<AddressEntity?> address = ValueNotifier<AddressEntity?>(null);

  Future<AddressEntity?> getAddress(String cep) async {
    isLoading.value = true;
    isError.value = "";

    try {
      final res = await getAddressUseCase.getAddress(cep);

      address.value = res;
      return res;
    } catch (e) {
      isError.value = e.toString();
    }

    return null;
  }

}
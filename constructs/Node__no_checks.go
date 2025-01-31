//go:build no_runtime_type_checking

package constructs

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Node) validateAddErrorParameters(message *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateAddInfoParameters(message *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateAddMetadataParameters(type_ *string, data interface{}) error {
	return nil
}

func (n *jsiiProxy_Node) validateAddValidationParameters(validation IValidation) error {
	return nil
}

func (n *jsiiProxy_Node) validateAddWarningParameters(message *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateApplyAspectParameters(aspect IAspect) error {
	return nil
}

func (n *jsiiProxy_Node) validateFindChildParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateSetContextParameters(key *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_Node) validateSynthesizeParameters(options *SynthesisOptions) error {
	return nil
}

func (n *jsiiProxy_Node) validateTryFindChildParameters(id *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateTryGetContextParameters(key *string) error {
	return nil
}

func (n *jsiiProxy_Node) validateTryRemoveChildParameters(childName *string) error {
	return nil
}

func validateNode_OfParameters(construct IConstruct) error {
	return nil
}

func validateNewNodeParameters(host Construct, scope IConstruct, id *string) error {
	return nil
}

